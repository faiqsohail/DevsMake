import { basePath } from '../path.config';

const createIdea = async (token, title, description) => {
    let result = null
    const response = await fetch(basePath + `/api/v1/ideas`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
            title: title,
            description: description
        })
    });

    if (response.ok) {
        result = await response.json()
    }
    
    return result
};

export default createIdea;
