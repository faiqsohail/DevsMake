import { useRouter } from 'next/router';
import * as React from 'react';

// TODO

const Search = () => {
    const router = useRouter()
    const { q } = router.query
  
    return <>
        <h1>Query: { decodeURIComponent(q) } </h1>
    </>
}

export default Search;