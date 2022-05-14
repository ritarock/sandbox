import type { NextPage } from 'next'
import ReactMarkdown from 'react-markdown'
import fs from 'fs'
import path from 'path'
import {CodeBlock} from '../components/codeBlock'

const Home: NextPage = ({markdown}: {markdown: string}) => {
  return (
    <div>
      <ReactMarkdown
        children={markdown}
        components={{
          code: CodeBlock
        }}
      />
    </div>
  )
}

export default Home

export const getStaticProps = async () => {
  const markdown = fs.readFileSync(
    path.join(process.cwd(), './pages/sample.md'),
    'utf8',
  )

  return {
    props: {
      markdown
    },
  }
}
