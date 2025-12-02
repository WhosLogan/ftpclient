import {h} from 'preact';
import {
    Button,
    Checkbox,
    Dialog, DialogActions, DialogContent, DialogContentText,
    DialogTitle,
    FormControlLabel,
    TextField,
} from "@mui/material";
import {useState} from "preact/compat";
import {Connect} from "../../wailsjs/go/main/App";
import { useLocation } from 'preact-iso';

export function HomePage(props: any) {
    const [useAnon, setUseAnon] = useState(true);
    const [server, setServer] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState<string|undefined>(undefined);
    const [loading, setLoading] = useState(false);
    const location = useLocation();

    const connect = async () => {
        setLoading(true);

        try {
            await Connect(server, username, password, useAnon);
            location.route('/browser');
        } catch (e) {
            setError(e as string);
        }

        setLoading(false);
    }

    return (
        <>
            <div className="flex flex-col h-screen items-center justify-center space-y-4 select-none">
                <h1 className="text-lg text-center">Welcome to FTP Client!</h1>
                <h2 className="text-center">To get started enter a server to connect to.</h2>

                <div className="flex flex-col md:w-1/3 w-5/6 gap-4">
                    <TextField label="Server" variant="outlined" disabled={loading} value={server} onChange={(e) => setServer(e.target.value)} />
                    <TextField label="Username" variant="outlined" disabled={useAnon || loading} value={username} onChange={(e) => setUsername(e.target.value)} />
                    <TextField label="Password" variant="outlined" disabled={useAnon || loading} value={password} onChange={(e) => setPassword(e.target.value)} />
                    <FormControlLabel
                        label="Use anonymous"
                        control={<Checkbox />}
                        checked={useAnon}
                        onChange={(_, checked) => setUseAnon(checked)}
                    />
                </div>

                <Button disabled={loading} variant="contained" onClick={connect}>Connect</Button>
            </div>


            <Dialog
                open={error !== undefined}
                onClose={() => setError(undefined)}
                aria-labelledby="alert-dialog-title"
                aria-describedby="alert-dialog-description"
                className="select-none"
            >
                <DialogTitle id="alert-dialog-title">
                    {"Unable to connect to server"}
                </DialogTitle>
                <DialogContent>
                    <DialogContentText id="alert-dialog-description">
                        {error}
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={() => setError(undefined)}>Close</Button>
                </DialogActions>
            </Dialog>
        </>
    )
}