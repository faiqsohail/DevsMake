import { Button, Container, Grid, TextField, Typography } from "@mui/material";
import CreateIcon from '@mui/icons-material/Create';
import LightbulbOutlinedIcon from '@mui/icons-material/LightbulbOutlined';
import { useEffect, useState } from "react";
import NavBar from "../../src/components/NavBar";
import useProfile from "../../src/hooks/useProfile";
import Cookies from "js-cookie";
import { basePath } from "../../src/path.config";
import Router from "next/router";
import createIdea from "../../src/api/createIdea";

const NewIdea = () => {
    const profile = useProfile();

    const sessionCookie = Cookies.get('sessionCookie')
    const isLoggedIn = sessionCookie != null

    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');

    useEffect(() => {
        if (typeof window !== "undefined") {
            setTitle(sessionStorage.getItem('idea_title') ?? '');
            setDescription(sessionStorage.getItem('idea_description') ?? '');
        }
    }, [])

    return (<>
        <NavBar profile={profile} />
        <section>
            <Container maxWidth="lg" sx={{ mt: 2 }}>
                <Grid container direction={"column"} spacing={2}>
                    <Grid item>
                        <Typography variant="h5" component="h3">
                            <LightbulbOutlinedIcon /> Idea Submission
                        </Typography>
                    </Grid>
                    <Grid item>
                        <TextField label="Title" id="title" value={title} onChange={(e) => {
                            setTitle(e.target.value)
                            sessionStorage.setItem("idea_title", e.target.value);
                        }} />
                    </Grid>
                    <Grid item>
                        <TextField multiline rows={10} value={description} fullWidth label="Description" id="description" placeholder="Describe your idea... You may answer questions like: What are the main features? Why is it useful? Why is it needed? Are there any alternatives? What platform would it be for?" onChange={(e) => {
                            setDescription(e.target.value)
                            sessionStorage.setItem("idea_description", e.target.value);
                        }} />
                    </Grid>
                    <Grid item>
                        <Button variant="contained" color="primary" startIcon={<CreateIcon />} onClick={() => {
                            if (title.length > 0 && description.length > 0) {
                                if (!isLoggedIn) {
                                    Router.push(basePath + `/api/v1/auth/login`)
                                } else {
                                    createIdea(sessionCookie, title, description).then((resp) => {
                                        if (resp != null) {
                                            sessionStorage.removeItem('idea_title')
                                            sessionStorage.removeItem('idea_description')

                                            Router.push(`/idea/${resp.uuid}`)
                                        }
                                    })
                                }
                            }
                        }}>Submit</Button>
                    </Grid>
                </Grid>
            </Container>
        </section>
    </>);
};

export default NewIdea;
