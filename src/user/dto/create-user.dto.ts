import { IsEmail, IsNotEmpty, IsString, Matches } from "class-validator";

export class CreateUserDTO {
    @IsEmail({}, {message: 'Please provide a valid email address'})
    @IsNotEmpty({message: 'Email is required'})
    email: string;

    @IsString({message: 'Username must be valid'})
    @IsNotEmpty({message: 'Username is required'})
    username: string;

    @IsString({message: 'Password must be valid'})
    @IsNotEmpty({message: 'Password is required'})
    @Matches(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{6,}/, {
      message: 'Password must contain at least one uppercase letter, one lowercase letter, one number and one special character, and be at least 6 characters long'
    })
    password: string;
}