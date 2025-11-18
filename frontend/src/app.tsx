import './App.css'
import {h} from 'preact';
import {LocationProvider, Router} from "preact-iso";
import {HomePage} from "./pages/home";

export function App(props: any) {
    return (
        <LocationProvider>
            <Router children={[
                <HomePage path="/" />
            ]} />
        </LocationProvider>
    )
}
