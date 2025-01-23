import { useNavigate } from "react-router-dom"
import React, { useEffect } from "react"

interface ProtectedRouteProps {
    children: React.ReactNode;
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({children}) => {
    const navigate = useNavigate();
    useEffect(() => {
        const token = localStorage.getItem("token");
        if(!token) {
            navigate("/login");
        }
    }, [navigate]
);

  return (
    <>
    {children}
    </>
)
}

export default ProtectedRoute
