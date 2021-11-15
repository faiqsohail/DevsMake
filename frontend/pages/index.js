import * as React from 'react';
import Container from '@mui/material/Container';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import { Button } from '@mui/material';
import LightbulbIcon from '@mui/icons-material/Lightbulb';
import NavBar from '../src/NavBar';
import { basePath } from '../path.config';

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
    <Container maxWidth="lg">
      <Box sx={{ my: 4, width: '100%',
        height: 300,
        backgroundColor: 'secondary.main' }}>
        <Typography variant="h4" component="h1" gutterBottom>
          Next.js v5 example
        </Typography>
      </Box>
      <Button variant="contained">Hey</Button>
    </Container>
    </>
  )
}
