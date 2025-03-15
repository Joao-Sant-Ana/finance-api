import { IsEmail, IsNotEmpty, IsString, Matches } from "class-validator";

export class RefreshUserTokenDTO {
    @IsEmail({}, {message: 'Please provide a valid email address'})
    @IsNotEmpty({message: 'Email is required'})
    email: string;

    @IsString({message: 'Id must be valid'})
    @IsNotEmpty({message: 'Id is required'})
    id: string;
}