import { Body, Controller, Post } from '@nestjs/common';
import { AuthService } from './auth.service';
import { SignupDto } from './dto/signup.dto';

@Controller('auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Post('/signup')
  async signUp(@Body() body: SignupDto) {
    return this.authService.signUp(body);
  }

  @Post('/login')
  async login(@Body() body: SignupDto) {
    return this.authService.login(body);
  }
}
