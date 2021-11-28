import { Avatar, Container, FormControl, Grid, InputLabel, NativeSelect, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Typography } from "@mui/material";
import LeaderboardIconOutlined from '@mui/icons-material/LeaderboardOutlined';
import NavBar from "../src/components/NavBar";
import useProfile from "../src/hooks/useProfile";
import { basePath } from "../src/path.config";
import Router, { useRouter } from "next/router";
import { SortByAlphaRounded } from "@mui/icons-material";

export async function getServerSideProps({ params, req }) {
    const sort = req.__NEXT_INIT_QUERY.sort ?? 'points'
    if (!['points', 'submissions', 'likes'].includes(sort)) {
        return {
            notFound: true,
        }
    }
    const response = await fetch(basePath + `/api/v1/profiles?sort=${sort}`)

    if (!response.ok) {
        return {
            notFound: true,
        }
    }

    let profiles = []
    profiles = await response.json()

    return {
        props: {
            profiles,
        }
    }
}

const Leaderboard = ({ profiles }) => {
    const profile = useProfile();

    const router = useRouter()
    const { sort } = router.query

    return (<>
        <NavBar profile={profile} />
        <section>
            <Container maxWidth="lg" sx={{ mt: 2 }}>
                <Grid container direction={"column"} spacing={2}>
                    <Grid item>
                        <Typography variant="h5" component="h3">
                            <LeaderboardIconOutlined /> Leaderboard
                        </Typography>
                    </Grid>
                    <Grid item>
                        <FormControl>
                            <InputLabel variant="standard" htmlFor="sort-native">
                                Sort
                            </InputLabel>
                            <NativeSelect
                                defaultValue={sort}
                                inputProps={{
                                    name: 'sort',
                                    id: 'sort-native',
                                }}
                                onChange={(e) => {
                                    Router.push({
                                        pathname: '/leaderboard',
                                        query: { sort: e.target.value }
                                    })
                                }}
                            >
                                <option value={'points'}>Points</option>
                                <option value={'submissions'}>Submissions</option>
                                <option value={'likes'}>Likes</option>
                            </NativeSelect>
                        </FormControl>
                    </Grid>
                    <Grid item>
                        <TableContainer component={Paper}>
                            <Table aria-label="leaderboard table">
                                <TableHead>
                                    <TableRow>
                                        <TableCell>Username</TableCell>
                                        <TableCell align="right">Points Earned</TableCell>
                                        <TableCell align="right">Submissions</TableCell>
                                        <TableCell align="right">Ratings</TableCell>
                                    </TableRow>
                                </TableHead>
                                <TableBody>
                                    {profiles.map((profile) => (
                                        <TableRow
                                            key={profile.username}
                                            sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                                        >
                                            <TableCell component="th" scope="row" onClick={() => {
                                                Router.push(`/profile/${profile.identifier}`)
                                            }}>
                                                <Grid container direction={"row"} spacing={2}>
                                                    <Grid item>
                                                        <Avatar sx={{ bgcolor: '#fff' }} src={profile.avatar_url} />
                                                    </Grid>
                                                    <Grid item>
                                                        <Typography variant="h5" component="h5">
                                                            {profile.username}({profile.identifier})
                                                        </Typography>
                                                    </Grid>
                                                </Grid>
                                            </TableCell>
                                            <TableCell align="right">{profile.points}</TableCell>
                                            <TableCell align="right">{profile.total_submissions}</TableCell>
                                            <TableCell align="right">{profile.total_ratings}</TableCell>
                                        </TableRow>
                                    ))}
                                </TableBody>
                            </Table>
                        </TableContainer>
                    </Grid>
                </Grid>
            </Container>
        </section>
    </>);
};

export default Leaderboard;
