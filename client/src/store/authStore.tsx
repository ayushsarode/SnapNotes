import {create} from 'zustand'

interface AuthState {
    token: string | null;
    setToken: (token: string) =>void;
    logout: () => void;
}

const useAuthStore = create<AuthState> ((set) => ({
    token: localStorage.getItem('token'),
    setToken: (token: string) => {
        localStorage.setItem('token', token);
        set({token});
    },
    logout: () => {
        localStorage.removeItem('token');
        set({token : null});
    },
}))


export default useAuthStore
