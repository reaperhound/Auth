import { IsLowercase, IsString, MinLength } from 'class-validator';

export class SignupDto {
  @IsString()
  @IsLowercase()
  username: string;

  @IsString()
  @MinLength(8)
  password: string;
}
