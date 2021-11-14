import * as React from 'react';
import Container from '@mui/material/Container';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';

import { Button } from '@mui/material';
import NavBar from '../src/NavBar';

export default function Home() {
  return (
    <>
    <NavBar />
    <Container maxWidth="sm">
      <Box sx={{ my: 4 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          Next.js v5 example
        </Typography>
      </Box>
      <Button variant="contained">Hey</Button>
    </Container>
    </>
  )
}
