import * as React from 'react';
import Container from '@mui/material/Container';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import { Button, Grid } from '@mui/material';
import LightbulbIcon from '@mui/icons-material/Lightbulb';
import NavBar from '../src/NavBar';
import { basePath } from '../path.config';
import IdeaCard from '../src/IdeaCard';

export async function getServerSideProps(context) {
  let profile = null;
  const token = context.req.cookies.sessionCookie ?? false;

  if (token) {
    const response = await fetch(basePath + `/api/v1/profile`, {
      headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
      },
    });

    if (response.ok) {
      profile = await response.json()
    }
  }

  return {
    props: {
      profile
    }
  }
}

export default function Home({profile}) {
  return (
    <>
      <NavBar profile={profile} />
      <section>
        <Container maxWidth="md">
          <Box py={8} textAlign="center">
            <Typography variant="h3" component="h2" gutterBottom={true}>DevsMake</Typography>
            <Typography variant="h5" color="textSecondary" paragraph={true}>A social platform bringing together users and developers to create amazing projects.</Typography>
            <Box mt={4}>
              <Button variant="contained" color="primary" startIcon={<LightbulbIcon />}>submit an idea</Button>
            </Box>
          </Box>
        </Container>
      </section>
      <section>
        <Container maxWidth="lg" sx={{mb: 10}}>
        <Grid container spacing={{ xs: 2, md: 3 }} columns={{ xs: 4, sm: 8, md: 12 }}>
          {Array.from(Array(6)).map((_, index) => (
            <Grid item xs={2} sm={4} md={4} key={index}>
              <IdeaCard />
            </Grid>
          ))}
        </Grid>
        </Container>
      </section>
    </>
  )
}
