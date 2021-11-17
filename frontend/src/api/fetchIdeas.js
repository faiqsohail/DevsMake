import { basePath } from '../path.config';

const fetchIdeas = async () => {
    let idea = [];

    const response = await fetch(basePath + `/api/v1/ideas`);

    if (response.ok) {
        idea = await response.json()
    }

    return idea
};

export default fetchIdeas;
