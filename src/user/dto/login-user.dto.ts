import { IsEmail, IsNotEmpty, IsString, Matches } from "class-validator";

export class LoginUserDTO {
    @IsEmail({}, {message: 'Please provide a valid email address'})
    @IsNotEmpty({message: 'Email is required'})
    email: string;

    @IsString({message: 'password must be valid'})
    @IsNotEmpty({message: 'password is required'})
    password: string;
}