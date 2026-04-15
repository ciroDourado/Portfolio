import { MetadataRoute } from 'next';
import { fetchAPI } from '@/lib/api';
import { Project } from '@/types';

export default async function sitemap(): Promise<MetadataRoute.Sitemap> {
    const baseUrl = process.env.SITE_URL || 'http://localhost:3000';

    let projects: Project[] = [];
    try {
        projects = await fetchAPI<Project[]>('/projects');
    } catch (error) {
        console.error("Erro ao gerar sitemap:", error);
    }

    const projectUrls = projects.map((project: Project) => ({
        url: `${baseUrl}/projects/${project.slug}`,
        lastModified: new Date,
        changeFrequency: 'monthly' as const,
        priority: 0.8,
    }));

    return [
        {
            url: baseUrl,
            lastModified: new Date,
            changeFrequency: 'weekly',
            priority: 1,
        },
        ...projectUrls,
    ];
}