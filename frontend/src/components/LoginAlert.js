import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle } from "@mui/material";
import Router from "next/router";
import { basePath } from "../path.config";


export default function LoginAlert({ open, setOpen }) {
    return (
        <Dialog
            open={open}
            onClose={() => setOpen(false)}
            aria-labelledby="alert-dialog-title"
            aria-describedby="alert-dialog-description"
        >
            <DialogTitle id="alert-dialog-title">
                {"Login required"}
            </DialogTitle>
            <DialogContent>
                <DialogContentText id="alert-dialog-description">
                    To perform this action you must login, would you like to login or sign-up now? It only takes a second!
                </DialogContentText>
            </DialogContent>
            <DialogActions>
                <Button onClick={() => setOpen(false)}>No</Button>
                <Button onClick={() => Router.push(basePath + `/api/v1/auth/login`)} autoFocus>
                    Yes
                </Button>
            </DialogActions>
        </Dialog>
    )
};
