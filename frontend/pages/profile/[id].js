import { basePath } from '../../src/path.config';
import { useRouter } from 'next/router'

export async function getStaticPaths() {
    return {
        paths: [],
        fallback: 'blocking'
    }
}
export async function getStaticProps({ params, req }) {
    const response = await fetch(basePath + `/api/v1/profile/${params.id}`)

    if (!response.ok) {
        return {
          notFound: true,
        }
    }

    const profile = await response.json()

    return {
        props: {
          profile,
        },
        revalidate: 10,
      }
  }

const ProfileById = ({profile}) => {
    const router = useRouter()
    const { id } = router.query
  
    return <>
        <h1>Todo Profile Id: {id}</h1>
        <h1>Username: {profile.username}</h1>
        <img src={profile.avatar_url} />
    </>
}

export default ProfileById;