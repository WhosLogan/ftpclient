import './App.css'
import {h} from 'preact';
import {LocationProvider, Router} from "preact-iso";
import {HomePage} from "./pages/home";
import { ThemeProvider, createTheme } from '@mui/material/styles';

const darkTheme = createTheme({
    palette: {
        mode: 'dark',
    },
});


export function App(props: any) {
    return (
        <ThemeProvider theme={darkTheme}>
            <LocationProvider>
                <Router children={[
                    <HomePage path="/" />
                ]} />
            </LocationProvider>
        </ThemeProvider>
    )
}
