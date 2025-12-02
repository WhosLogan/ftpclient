import {useEffect, useState} from "preact/hooks";
import {ChangeDirectory, Close, ListDirectory} from "../../wailsjs/go/main/App";
import {ftpclient} from "../../wailsjs/go/models";
import FtpEntry = ftpclient.FtpEntry;
import {Button, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow} from "@mui/material";
import {useLocation} from "preact-iso";

export function BrowserPage(props: any) {
    const [dir, setDir] = useState<FtpEntry[]>([]);
    const location = useLocation();

    useEffect(() => {
        refresh().then();
    }, []);

    const disconnect = async () => {
        await Close();
        location.route('/');
    }

    const cd = async (path: string) => {
        try {
            await ChangeDirectory(path);
        } catch (e) {
            console.error(e);
        }
        await refresh()
        console.log(dir);
    }

    const refresh = async () => {
        try {
            setDir(await ListDirectory());
        } catch (e) {
            console.error(e);
        }
    }

    return (
        <>
            <div className="flex flex-col gap-4" style={{height: "calc(100vh - 128px)"}}>
                <TableContainer component={Paper}>
                    <Table sx={{ minWidth: 650 }} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell>Name</TableCell>
                                <TableCell align="right">Directory</TableCell>
                                <TableCell align="right">Permissions</TableCell>
                                <TableCell align="right">Size</TableCell>
                                <TableCell align="right">Actions</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            <TableRow
                                key={'..'}
                                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                            >
                                <TableCell component="th" scope="row">
                                    ..
                                </TableCell>
                                <TableCell align="right">True</TableCell>
                                <TableCell align="right">N/A</TableCell>
                                <TableCell align="right">N/A</TableCell>
                                <TableCell align="right">
                                    <Button variant="contained" onClick={() => cd("..")}>Open</Button>
                                </TableCell>
                            </TableRow>

                            {dir.map((entry) => (
                                <TableRow
                                    key={entry.Name}
                                    sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                                >
                                    <TableCell component="th" scope="row">
                                        {entry.Name}
                                    </TableCell>
                                    <TableCell align="right">{entry.IsDir ? "True" : "False"}</TableCell>
                                    <TableCell align="right">{entry.Perms}</TableCell>
                                    <TableCell align="right">{entry.IsDir ? "N/A" : entry.Size}</TableCell>
                                    <TableCell align="right">
                                        {entry.IsDir ? <Button variant="contained" onClick={() => cd(entry.Name)}>Open</Button> :
                                            <Button variant="contained">Download</Button>}
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            </div>

            <Button variant="contained" onClick={disconnect}>Disconnect</Button>
            <Button variant="contained" onClick={refresh}>Refresh</Button>
        </>
    )
}