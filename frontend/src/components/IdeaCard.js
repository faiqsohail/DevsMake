import { Badge, Card, CardActionArea, CardActions, CardContent, Grid, Typography } from "@mui/material";
import LikeIcon from '@mui/icons-material/ThumbUpAlt';
import DislikeIcon from '@mui/icons-material/ThumbDownAlt';
import CodeIcon from '@mui/icons-material/Code';
import Router from 'next/router';

export default function IdeaCard({ uuid, title, description, likes, dislikes, submissions }) {
  return (
    <Card sx={{ maxWidth: 345 }} onClick={() => Router.push(`/idea/${uuid}`)}>
      <CardActionArea >
        <CardContent sx={{ minHeight: 200, maxHeight: 200 }}>
          <Typography gutterBottom variant="h5" component="div">
            {title}
          </Typography>
          <Typography variant="body2" color="text.secondary">
            {description}
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
            <Badge badgeContent={likes} color="primary">
              <LikeIcon color="action" />
            </Badge>
            <Badge badgeContent={dislikes} color="primary">
              <DislikeIcon color="action" />
            </Badge>
          </Grid>
          <Grid item sx={{ paddingRight: '10px' }}>
            <Badge badgeContent={submissions} color="primary">
              <CodeIcon color="action" />
            </Badge>
          </Grid>
        </Grid>
      </CardActions>
    </Card>
  );
}
