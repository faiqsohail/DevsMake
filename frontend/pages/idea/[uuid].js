import Router, { useRouter } from "next/router"
import NavBar from "../../src/components/NavBar"
import { basePath } from "../../src/path.config"
import useProfile from '../../src/hooks/useProfile';
import LikeIcon from '@mui/icons-material/ThumbUpAlt';
import DislikeIcon from '@mui/icons-material/ThumbDownAlt';
import CodeIcon from '@mui/icons-material/Code';
import SendIcon from '@mui/icons-material/Send';
import { Avatar, Badge, Card, CardActionArea, CardActions, CardContent, CardHeader, Container, Grid, TextField, Typography } from "@mui/material";

export async function getStaticPaths() {
    return {
        paths: [],
        fallback: 'blocking'
    }
}
export async function getStaticProps({ params, req }) {
    const response = await fetch(basePath + `/api/v1/ideas/${params.uuid}`)

    if (!response.ok) {
        return {
            notFound: true,
        }
    }

    const ideaPost = await response.json()
    const ideaComments = await (await fetch(basePath + `/api/v1/ideas/${params.uuid}/comments`)).json()
    const authorProfile = await (await fetch(basePath + `/api/v1/profile/${ideaPost.author_id}`)).json()

    for (const comment of ideaComments) {
        const commenter = await (await fetch(basePath + `/api/v1/profile/${comment.author_id}`)).json()
        Object.assign(comment, { profile: commenter })
    }

    return {
        props: {
            ideaPost,
            ideaComments,
            authorProfile
        },
        revalidate: 30,
    }
}

const IdeaByUUID = ({ ideaPost, ideaComments, authorProfile }) => {
    const router = useRouter()
    const profile = useProfile()

    const { uuid } = router.query

    return (
        <>
            <NavBar profile={profile} />
            <section>
                <Container maxWidth="lg" sx={{ mt: 2 }}>

                    <Grid container spacing={2} columns={16} sx={{ mt: 2, backgroundColor: '#fff', }}>
                        <Grid item xs={8}>
                            <Card>
                                <CardHeader avatar={
                                    <Avatar sx={{ bgcolor: '#fff' }} src={authorProfile.avatar_url} />
                                }
                                    title={`Author: ${authorProfile.username}(${authorProfile.identifier})`}
                                    onClick={() => Router.push(`/profile/${authorProfile.identifier}`)} />
                                <CardActionArea >
                                    <CardContent sx={{ minHeight: 200, maxHeight: 200 }}>
                                        <Typography gutterBottom variant="h5" component="div">
                                            {ideaPost.title}
                                        </Typography>
                                        <Typography variant="body2" color="text.secondary">
                                            {ideaPost.description}
                                        </Typography>
                                    </CardContent>
                                </CardActionArea>
                                <CardActions>
                                    <Grid
                                        container
                                        direction="row"
                                        justifyContent="space-between"
                                        alignItems="baseline"
                                    >
                                        <Grid item>
                                            <Badge badgeContent={ideaPost.likes} color="primary">
                                                <LikeIcon color="action" />
                                            </Badge>
                                            <Badge badgeContent={ideaPost.dislikes} color="primary">
                                                <DislikeIcon color="action" />
                                            </Badge>
                                        </Grid>
                                        <Grid item sx={{ paddingRight: '10px' }}>
                                            <Badge badgeContent={ideaPost.submissions} color="primary">
                                                <CodeIcon color="action" />
                                            </Badge>
                                        </Grid>
                                    </Grid>
                                </CardActions>
                            </Card>
                        </Grid>
                        <Grid item xs={8}>
                            <TextField fullWidth label="Leave a comment" id="comment" InputProps={{ endAdornment: <SendIcon /> }} />
                            {ideaComments.length > 0 ? ideaComments.map(comment => (
                                <Card key={comment.uuid}>
                                    <CardHeader avatar={
                                        <Avatar sx={{ bgcolor: '#fff' }} src={comment.profile.avatar_url} />
                                    }
                                        title={`Author: ${comment.profile.username}(${comment.profile.identifier})`}
                                        onClick={() => Router.push(`/profile/${comment.profile.identifier}`)} />
                                    <CardContent>
                                        {comment.comment}
                                    </CardContent>
                                </Card>
                            )) : <p>No comments posted</p>}
                        </Grid>
                    </Grid>
                </Container>
            </section>
        </>
    );

}

export default IdeaByUUID;