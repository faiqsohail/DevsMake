import { basePath } from '../../src/path.config';
import { useRouter } from 'next/router'
import NavBar from '../../src/components/NavBar';
import Image from 'next/image'
import { Avatar, Container, Grid, Typography } from '@mui/material';
import { Box } from '@mui/system';
import useProfile from '../../src/hooks/useProfile';

export async function getStaticPaths() {
    return {
        paths: [],
        fallback: 'blocking'
    }
}
export async function getStaticProps({ params, req }) {
    const response = await fetch(basePath + `/api/v1/profile/${params.id}`)

    if (!response.ok) {
        return {
            notFound: true,
        }
    }

    const profile = await response.json()

    return {
        props: {
            profile,
        },
        revalidate: 10,
    }
}

const ProfileById = ({ profile }) => {
    const router = useRouter()
    const { id } = router.query

    const currentProfile = useProfile();

    const post_badges = [1, 10, 50, 100, 500];

    return (
        <>
            <NavBar profile={currentProfile} />
            <section>
                <Container maxWidth="md" sx={{ mt: 2 }}>

                    <Grid container spacing={2} columns={16} sx={{ mt: 2, backgroundColor: '#fff', }}>
                        <Grid item xs={8}>
                            <Box>
                                <center>
                                    <Avatar
                                        alt={profile.username}
                                        src={profile.avatar_url}
                                        sx={{ width: 256, height: 246, border: '0.5px solid lightgray', backgroundColor: '#fff' }}
                                    />
                                    <Typography variant="h3" color="textSecondary">{profile.username}({id})</Typography>
                                    <Typography variant="h7" color="textSecondary">Ideas posted: {profile.total_posts}</Typography><br />
                                    <Typography variant="h7" color="textSecondary">Submissions:  {profile.total_submissions}</Typography>
                                </center>
                            </Box>
                        </Grid>
                        <Grid item xs={8}>
                            <Typography variant="h5" color="textSecondary">Badges</Typography>
                            {profile.total_posts == 0 ? <p>Make a idea post to earn a badge!</p> : post_badges.map((badge_id) => {
                                if (profile.total_posts >= badge_id) {
                                    return <Image width={100} height={128} id={badge_id} alt={'post badge'} src={`/badges/post/${badge_id}.png`} />
                                }
                            })}
                        </Grid>
                    </Grid>
                </Container>
            </section>
        </>
    )
}

export default ProfileById;