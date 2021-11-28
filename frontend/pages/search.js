import { Container, Grid, Typography } from '@mui/material';
import SearchOutlinedIcon from '@mui/icons-material/SearchOutlined';
import { useRouter } from 'next/router';
import * as React from 'react';
import NavBar from '../src/components/NavBar';
import useProfile from '../src/hooks/useProfile';
import IdeaCard from '../src/components/IdeaCard';
import fetchIdeas from '../src/api/fetchIdeas';

export async function getServerSideProps({ params, req }) {
    const query = req.__NEXT_INIT_QUERY.q
    if (!query) {
        return {
            notFound: true,
        }
    }
    const ideas = await fetchIdeas(query);

    return {
        props: {
            ideas,
        }
    }
}

const Search = ({ ideas }) => {
    const router = useRouter()
    const { q } = router.query

    const profile = useProfile();

    return (<>
        <NavBar profile={profile} />
        <section>
            <Container maxWidth="lg" sx={{ mt: 2 }}>
                <Grid container direction={"column"} spacing={2}>
                    <Grid item>
                        <Typography variant="h5" component="h3">
                            <SearchOutlinedIcon /> Search Result: {decodeURIComponent(q)}
                        </Typography>
                    </Grid>
                    <Grid item>
                        {ideas.length > 0 ? 
                        <Grid container spacing={{ xs: 2, md: 3 }} columns={{ xs: 4, sm: 8, md: 12 }}>
                            {ideas.map((idea) => (
                                <Grid item xs={2} sm={4} md={4} key={idea.uuid}>
                                    <IdeaCard {...idea} />
                                </Grid>
                            ))}
                        </Grid>
                        :  <p>No idea posts found matching the search criteria</p> }
                    </Grid>
                </Grid>
            </Container>
        </section>
    </>)
}

export default Search;
