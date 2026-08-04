import {
  BadRequestException,
  Injectable,
  UnauthorizedException,
} from '@nestjs/common';
import { SignupDto } from './dto/signup.dto';
import { FileService } from './file.service';
import * as argon2 from 'argon2';
import { JwtService } from '@nestjs/jwt';
import { ConfigService } from '@nestjs/config';
import { RefreshDTO } from './dto/refresh.dto';

@Injectable()
export class AuthService {
  constructor(
    private readonly fileService: FileService,
    private readonly jwtService: JwtService,
    private readonly configService: ConfigService,
  ) {}
  async signUp(body: SignupDto) {
    const users = await this.fileService.readFile();
    const exists = users?.some((u) => u.username === body.username);

    if (exists) {
      throw new BadRequestException(
        `username ${body.username} is already taken`,
      );
    }

    const hashedPassword = await argon2.hash(body.password);

    const res = await this.fileService.appendFile({
      ...body,
      password: hashedPassword,
    });

    return {
      username: res.username,
      message: 'Signup successful',
    };
  }

  async login(body: SignupDto) {
    const users = await this.fileService.readFile();
    const user = users?.find((u) => u.username === body.username);

    if (!user) {
      throw new UnauthorizedException('Invalid credentials');
    }

    const isMatch = await argon2.verify(user.password, body.password);

    if (!isMatch) {
      throw new UnauthorizedException('Invalid credentials');
    }

    const payload = { sub: user.username, username: user.username };
    const accessToken = await this.jwtService.signAsync(payload, {
      secret: this.configService.get<string>('JWT_SECRET'),
      expiresIn: '15m',
    });

    const refreshToken = await this.jwtService.signAsync(payload, {
      secret: this.configService.get<string>('JWT_REFRESH_SECRET'),
      expiresIn: '7d',
    });

    const hashedRefreshToken = await argon2.hash(refreshToken);
    await this.fileService.updateUser(user.username, {
      refreshToken: hashedRefreshToken,
    });

    return { message: 'Login successful', accessToken, refreshToken };
  }

  async refresh(body: RefreshDTO) {
    let decoded: {
      sub: string;
      username: string;
    };
    try {
      decoded = await this.jwtService.verifyAsync(body.refreshToken, {
        secret: this.configService.get<string>('JWT_REFRESH_SECRET'),
      });
    } catch {
      throw new UnauthorizedException('Invalid or expired refresh token');
    }

    const users = await this.fileService.readFile();
    const user = users?.find((u) => u.username === decoded.username);

    if (!user || !user.refreshToken) {
      throw new UnauthorizedException('Invalid refresh token');
    }

    const isValid = await argon2.verify(user.refreshToken, body.refreshToken);
    if (!isValid) throw new UnauthorizedException('Invalid refresh token');

    const payload = { sub: user.username, username: user.username };
    const accessToken = await this.jwtService.signAsync(payload, {
      secret: this.configService.get<string>('JWT_SECRET'),
      expiresIn: '15m',
    });

    const refreshToken = await this.jwtService.signAsync(payload, {
      secret: this.configService.get<string>('JWT_REFRESH_SECRET'),
      expiresIn: '7d',
    });

    const hashedRefreshToken = await argon2.hash(refreshToken);
    await this.fileService.updateUser(user.username, {
      refreshToken: hashedRefreshToken,
    });

    return { accessToken, refreshToken };
  }
}
