import { basePath } from '../path.config';

const commentIdea = async (token, uuid, comment) => {
    let result = null
    const response = await fetch(basePath + `/api/v1/ideas/${uuid}/comments`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
            comment: comment
        })
    });

    if (response.ok) {
        result = await response.json()
    }
    
    return result
};

export default commentIdea;
