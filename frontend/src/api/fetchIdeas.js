import { basePath } from '../path.config';

const fetchIdeas = async (query) => {
    let idea = [];
    let endpoint = '/api/v1/ideas'
    if (query) {
        endpoint = endpoint + `?query=${query}`
    }

    const response = await fetch(basePath + endpoint);

    if (response.ok) {
        idea = await response.json()
    }

    return idea
};

export default fetchIdeas;
