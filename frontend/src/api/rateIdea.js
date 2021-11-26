import { basePath } from '../path.config';

const rateIdea = async (token, uuid, rating) => {
    let result = null
    const response = await fetch(basePath + `/api/v1/ideas/${uuid}/rate`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
            rating: rating
        })
    });

    if (response.ok) {
        result = await response.json()
    }

    return result
};

export default rateIdea;
