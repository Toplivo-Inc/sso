import { createRoot } from 'react-dom/client'
import './index.css'
import { RouterProvider } from 'react-router'
import { ThemeProvider } from '@/components/theme-provider'
import { router } from '@/app/router'

const root = document.getElementById('root')!


createRoot(root).render(
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
        <RouterProvider router={router} />
    </ThemeProvider>
)

// function Redirect() {
//     const navigate = useNavigate();
//     useEffect(() => {
//         navigate('/console', { replace: true })
//     });
//     return (<p>Redirecting</p>)
// }
