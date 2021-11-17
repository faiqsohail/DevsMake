import { useEffect, useState } from "react";
import Cookies from 'js-cookie'
import fetchProfile from "../api/fetchProfile";

const useProfile = () => {
    const [profile, setProfile] = useState(null);

    useEffect(() => {
        fetchProfile(Cookies.get('sessionCookie')).then(response => {
            setProfile(response);
        })
    }, [])

    return profile;
}

export default useProfile;
