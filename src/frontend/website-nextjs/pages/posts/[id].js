import Layout from '../../components/layout'
import Head from 'next/head'
import { getPostData } from '../../libs/posts'
import utilStyles from '../../styles/utils.module.css'
import Date from '../../components/date'

export default function Post({ postData }) {
    return (
        <Layout>
            <Head>
                <title>{postData.title}</title>
            </Head>
            <article>
                <h1 className={utilStyles.headingXl}>{postData.title}</h1>
                <div className={utilStyles.lightText}>
                    <Date dateString={postData.date} />
                </div>
                <div dangerouslySetInnerHTML={{ __html: postData.contentHtml }} />
            </article>
        </Layout>
    )
}

export async function getServerSideProps(context) {
    const { id } = context.query;
    console.log("id: " + id)
    const postData = await getPostData(id)
    return {
        props: {
            postData
        }
    }
}