import * as React from 'react';
import Container from '@mui/material/Container';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import { Button, Grid } from '@mui/material';
import LightbulbIcon from '@mui/icons-material/Lightbulb';
import NavBar from '../src/components/NavBar';
import IdeaCard from '../src/components/IdeaCard';
import fetchIdeas from '../src/api/fetchIdeas';
import useProfile from '../src/hooks/useProfile';
import Router from 'next/router';

export async function getStaticProps({ params, req }) {
  const ideas = await fetchIdeas();

  return {
      props: {
        ideas,
      },
      revalidate: 10,
    }
}

const Home = ({ ideas }) => {
  const profile = useProfile();

  return (
    <>
      <NavBar profile={profile} />
      <section>
        <Container maxWidth="md">
          <Box py={8} textAlign="center">
            <Typography variant="h3" component="h2" gutterBottom={true}>DevsMake</Typography>
            <Typography variant="h5" color="textSecondary" paragraph={true}>A social platform bringing together users and developers to create amazing projects.</Typography>
            <Box mt={4}>
              <Button variant="contained" color="primary" startIcon={<LightbulbIcon />} onClick={() => Router.push(`/idea/new`)}>submit an idea</Button>
            </Box>
          </Box>
        </Container>
      </section>
      <section>
        <Container maxWidth="lg" sx={{ mb: 10 }}>
          <Grid container spacing={{ xs: 2, md: 3 }} columns={{ xs: 4, sm: 8, md: 12 }}>
            {ideas.map((idea) => (
              <Grid item xs={2} sm={4} md={4} key={idea.uuid}>
                <IdeaCard {...idea} />
              </Grid>
            ))}
          </Grid>
        </Container>
      </section>
    </>
  )
}

export default Home;
