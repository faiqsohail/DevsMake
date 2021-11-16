import { useState, useEffect } from 'react';
import { basePath } from '../path.config';

const fetchProfile = async (token) => {
    let profile = null

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

    return profile
};

export default fetchProfile;