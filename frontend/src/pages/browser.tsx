import {useEffect} from "preact/hooks";
import {useState} from "preact/compat";
import {ListDirectory} from "../../wailsjs/go/main/App";
import {ftpclient} from "../../wailsjs/go/models";
import FtpEntry = ftpclient.FtpEntry;
import {Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow} from "@mui/material";

export function BrowserPage(props: any) {
    const [dir, setDir] = useState<FtpEntry[]>([]);

    useEffect(() => {
        console.log("ddasdasd")
        ListDirectory().then(a => {
            setDir(a);
        });
    }, []);

    return (
        <>
            <TableContainer component={Paper}>
                <Table sx={{ minWidth: 650 }} aria-label="simple table">
                    <TableHead>
                        <TableRow>
                            <TableCell>Name</TableCell>
                            <TableCell align="right">Directory</TableCell>
                            <TableCell align="right">Permissions</TableCell>
                            <TableCell align="right">Size</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
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
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </>
    )
}