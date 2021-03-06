import { serialize } from 'next-mdx-remote/serialize';
import { GetStaticProps, GetStaticPaths } from 'next';
import { MDXRemote, MDXRemoteSerializeResult } from 'next-mdx-remote';
import { IPost } from '../../types/post';
import Prerequisites from '../../components/Prerequisites';
import Stacks from '../../components/Stacks';
import { useMdxComponentsContext } from '../../context/mdxContext';
import { useEffect } from 'react';
import { ParsedUrlQuery } from 'querystring';
import { getAllFiles, getFileBySlug } from '../../lib/mdx';

type Props = {
		source: MDXRemoteSerializeResult;
		frontMatter: Omit<IPost, "slug">;
}

const components = {
	Prerequisites,
	Stacks
};

const PostPage: React.FC<Props> = ({ source, frontMatter} : Props) => {
	// get setters
	const { setPrerequisites, setStacks } = useMdxComponentsContext();

	useEffect(() => {
		setPrerequisites(frontMatter.prerequisites);
		setStacks(frontMatter.stacks);
	}, [
		setPrerequisites,
		setStacks,
		frontMatter.prerequisites,
		frontMatter.stacks
	]);

	return (
		<div>
			<article className="prose prose-green dark:text-slate-300">
				<h1 className="dark:text-slate-300">{frontMatter.title}</h1>

				<p>{frontMatter.description}</p>

				<MDXRemote components={components} {...source} />
			</article>
		</div>
	)
}

export default PostPage;

interface Iparams extends ParsedUrlQuery {
		slug: string;
};

export const getStaticProps:GetStaticProps = async (context) => {
	const { slug } = context.params as Iparams;

	const { content, data } = getFileBySlug(slug);

	const mdxSource = await serialize(content, { scope: data });
	return {
		props: {
			source: mdxSource,
			frontMatter: data
		}
	};
}

export const getStaticPaths: GetStaticPaths = () => {
	const posts = getAllFiles(['slug']);

	const paths = posts.map((post) => ({
			params: {
				slug: post.slug
			}
	}));

	return {
		paths,
		fallback: false
	};
}
