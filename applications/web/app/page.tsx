import { fetchAPI } from '@/lib/api';
import {Experience, Project} from '@/types';
import Link from 'next/link';
import ExperienceTimeline from "@/app/components/ExperienceTimeline";

export const dynamic = 'force-dynamic';

export default async function Home() {

  let projects: Project[] = [];
  try {
    projects = await fetchAPI<Project[]>('/projects');
  } catch (error) {
    console.error("Falha ao carregar projetos:", error);
  }

  let experiences: Experience[] = [];
  try {
    experiences = await fetchAPI<Experience[]>('/experiences');
  } catch (error) {
    console.error("Falha ao carregar experiências:", error);
  }

  return (
      <main className="min-h-screen p-8 md:p-24 bg-gray-50 dark:bg-gray-900">
        <div className="max-w-7xl mx-auto">
          <h1 className="text-4xl font-bold mb-8 text-gray-900 dark:text-white">
            Meus Projetos <span className="text-blue-500">(Next.js + Go)</span>
          </h1>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {projects.map((project) => (
                <div
                    key={project.id}
                    className="group bg-white dark:bg-gray-800 rounded-xl shadow-sm hover:shadow-md transition-all overflow-hidden border border-gray-200 dark:border-gray-700"
                >
                  <Link href={`/projects/${project.slug}`} className="block">
                    <div className="relative h-48 w-full bg-gray-200">
                      <img
                          src={project.thumbnail_url}
                          alt={project.title}
                          className="object-cover w-full h-full group-hover:scale-105 transition-transform duration-300"
                      />
                    </div>
                  </Link>

                  <div className="p-6">
                    <Link href={`/projects/${project.slug}`}>
                      <h3 className="text-xl font-bold text-gray-900 dark:text-white mb-2">
                        {project.title}
                      </h3>
                    </Link>
                    <p className="text-gray-500 dark:text-gray-400 text-sm line-clamp-3 mb-4">
                      {project.short_description}
                    </p>

                    <div className="flex flex-wrap gap-2">
                      {project.skills?.map((skill) => (
                          <span
                              key={skill.id}
                              className="px-2 py-1 text-xs font-medium bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200 rounded-full"
                          >
                      {skill.name}
                    </span>
                      ))}
                    </div>
                  </div>
                </div>
            ))}
          </div>

          {projects.length === 0 && (
              <p className="text-gray-500 mt-10">Nenhum projeto encontrado.</p>
          )}
        </div>

        <section className="mt-24 mb-12">
          <h2 className="text-3xl font-bold mb-12 text-gray-900 dark:text-white">
            Minha Jornada
          </h2>
          <ExperienceTimeline experiences={experiences} />
        </section>
      </main>
  );
}