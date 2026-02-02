import { fetchAPI } from '@/lib/api';
import { Project } from '@/types';
import Link from 'next/link';
import { notFound } from 'next/navigation';

export const dynamic = 'force-dynamic';

interface PageProps {
    params: Promise<{ slug: string }>;
}

export default async function ProjectDetails({ params }: PageProps) {
    const { slug } = await params;

    let project: Project | null = null;

    try {
        project = await fetchAPI<Project>(`/projects/${slug}`);
    } catch (error) {
        console.error(`Erro ao buscar projeto ${slug}:`, error);
    }

    if (!project) {
        notFound();
    }

    return (
        <main className="min-h-screen bg-gray-50 dark:bg-gray-900 py-12 px-4 sm:px-6 lg:px-8">
            <div className="max-w-4xl mx-auto">
                <Link
                    href="/"
                    className="inline-flex items-center text-sm font-medium text-gray-500 hover:text-gray-900 dark:text-gray-400 dark:hover:text-white mb-8 transition-colors"
                >
                    ← Voltar para Projetos
                </Link>

                <article className="bg-white dark:bg-gray-800 rounded-2xl shadow-xl overflow-hidden border border-gray-100 dark:border-gray-700">

                    <div className="relative h-64 md:h-96 w-full bg-gray-200">
                        <img
                            src={project.thumbnail_url}
                            alt={project.title}
                            className="w-full h-full object-cover"
                        />
                    </div>

                    <div className="p-8 md:p-12">
                        <div className="flex flex-col md:flex-row md:items-start md:justify-between gap-6 mb-8">
                            <div>
                                <h1 className="text-3xl md:text-4xl font-bold text-gray-900 dark:text-white mb-4">
                                    {project.title}
                                </h1>

                                <div className="flex flex-wrap gap-2">
                                    {project.skills?.map((skill) => (
                                        <span
                                            key={skill.id}
                                            className="px-3 py-1 text-sm font-medium bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200 rounded-full"
                                        >
                                          {skill.name}
                                        </span>
                                    ))}
                                </div>
                            </div>

                            <div className="flex gap-4 shrink-0">
                                {project.repository_url && (
                                    <a
                                        href={project.repository_url}
                                        target="_blank"
                                        rel="noopener noreferrer"
                                        className="px-6 py-2.5 bg-gray-900 text-white hover:bg-gray-800 rounded-lg font-medium transition-colors flex items-center gap-2"
                                    >
                                        <span>GitHub</span>
                                    </a>
                                )}

                                {project.live_url && (
                                    <a
                                        href={project.live_url}
                                        target="_blank"
                                        rel="noopener noreferrer"
                                        className="px-6 py-2.5 bg-blue-600 text-white hover:bg-blue-700 rounded-lg font-medium transition-colors"
                                    >
                                        Ver Projeto
                                    </a>
                                )}
                            </div>
                        </div>

                        <hr className="border-gray-200 dark:border-gray-700 my-8" />

                        <div className="prose prose-lg dark:prose-invert max-w-none text-gray-600 dark:text-gray-300">
                            <p className="whitespace-pre-line leading-relaxed">
                                {project.content || project.short_description}
                            </p>
                        </div>
                    </div>
                </article>
            </div>
        </main>
    );
}