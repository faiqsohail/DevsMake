import { useRouter } from 'next/router'

const ProfileById = () => {
    const router = useRouter()
    const { id } = router.query
  
    return <h1>Todo Profile Id: {id}</h1>
}

export default ProfileById;