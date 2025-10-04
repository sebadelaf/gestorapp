export interface User {
    id?: string;
    name: string;
    email: string;
    password: string;
}

//Dto para creacion de usuario
export interface CreateUserDto {
    name: string;
    email: string;
    password: string;
}