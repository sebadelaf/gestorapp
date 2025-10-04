import type { User, CreateUserDto } from '../types/user';

//Crear un usuario
export const createUser = async(userData: CreateUserDto): Promise<User>=>{
    const config = useRuntimeConfig();
    return await $fetch('/user/new',
        {method:'POST',
            body:userData,
            baseURL:config.public.apiBaseUrl
        }
    );
};
//obtener todos los usuarios
export const getUsers= async(): Promise<User[]>=>{
    const config = useRuntimeConfig();
    return await $fetch('/user/',
        {
            method:'GET',
            baseURL:config.public.apiBaseUrl
        }
    );
};
//obtener un usuario por id

export const getUserById= async(id:String):Promise<User>=>{
    const config = useRuntimeConfig();
    return await $fetch('/user/${id}',
        {
            method:'GET',
            baseURL:config.public.apiBaseUrl
        }
    );
};

//actualizar un usuario

//eliminar un usuario