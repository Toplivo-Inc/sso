import { LoginForm } from '@/pages/forms/login';
import { RegisterForm } from '@/pages/forms/register';
import { RouteObject } from 'react-router';

export const formRoutes: RouteObject[] = [
    {
        path: "/login",
        element: <LoginForm />,
    },
    {
        path: "/register",
        element: <RegisterForm />,
    },
]
