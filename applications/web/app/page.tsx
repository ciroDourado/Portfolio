import { fetchAPI } from '@/lib/api';
import {Experience, Profile, Project} from '@/types';
import Link from 'next/link';
import ExperienceTimeline from "@/app/components/ExperienceTimeline";
import ContactForm from '@/components/ContactForm';

export const dynamic = 'force-dynamic';

export const metadata = {
  title: 'Portfolio | Ciro Oliveira',
};

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

  let profile: Profile|null = null;
  try {
    profile = await fetchAPI<Profile>('/profile');
  } catch (error) {
    console.error("Falha ao carregar perfil:", error);
  }

  return (
      <main className="min-h-screen bg-gray-50 dark:bg-gray-900">
        {profile && (<section className="bg-white dark:bg-gray-800 py-20">
          <div className="max-w-7xl mx-auto px-4 flex flex-col md:flex-row items-center gap-12">

            <div className="relative w-48 h-48 md:w-64 md:h-64 shrink-0">
              {profile.avatar_url && (<img
                  src={profile.avatar_url}
                  alt={profile.name}
                  className="rounded-full object-cover w-full h-full shadow-2xl border-4 border-white dark:border-gray-700"
              />)}
            </div>

            <div className="text-center md:text-left">
              <h1 className="text-4xl md:text-6xl font-extrabold text-gray-900 dark:text-white mb-4">
                {profile.name}
              </h1>
              <p className="text-xl md:text-2xl text-blue-600 font-medium mb-6">
                {profile.headline}
              </p>
              <div
                  className="prose prose-lg dark:prose-invert text-gray-600 dark:text-gray-400 max-w-2xl text-lg mb-8 leading-relaxed whitespace-pre-line"
                  dangerouslySetInnerHTML={{ __html: profile.bio }}
              />

              <div className="flex flex-wrap justify-center md:justify-start gap-4">
                {profile.cv_url && (
                    <a href={profile.cv_url} target="_blank"
                       className="px-6 py-3 bg-blue-600 text-white rounded-lg font-bold hover:bg-blue-700 transition">
                      Download CV
                    </a>
                )}
                {profile.github_url && (
                    <a href={profile.github_url} target="_blank"
                       className="px-6 py-3 bg-gray-900 text-white rounded-lg font-bold hover:bg-gray-700 transition">
                      GitHub
                    </a>
                )}
                {profile.linkedin_url && (
                    <a href={profile.linkedin_url} target="_blank"
                       className="px-6 py-3 bg-blue-800 text-white rounded-lg font-bold hover:bg-blue-900 transition">
                      LinkedIn
                    </a>
                )}
              </div>
            </div>
          </div>
        </section>)}

        <div className="max-w-7xl mx-auto px-4 md:px-8 py-12">
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

          <section className="mt-24 mb-12">
            <h2 className="text-3xl font-bold mb-12 text-gray-900 dark:text-white">
              Minha Jornada
            </h2>
            <ExperienceTimeline experiences={experiences} />
          </section>
        </div>

        <section className="py-20 bg-gray-50 dark:bg-gray-900" id="contact">
          <div className="max-w-7xl mx-auto px-4 md:px-8">
            <div className="grid md:grid-cols-2 gap-12 items-center">

              <div>
                <h2 className="text-3xl md:text-4xl font-bold text-gray-900 dark:text-white mb-6">
                  Interessado no meu trabalho?
                </h2>
                <p className="text-lg text-gray-600 dark:text-gray-400 mb-8 leading-relaxed">
                  Estou sempre aberto a novas oportunidades e desafios.
                  Se você tem uma ideia de projeto ou quer discutir sobre tecnologia,
                  minha caixa de entrada está aberta.
                </p>

                <div className="space-y-4">
                  <div className="flex items-center gap-4 text-gray-700 dark:text-gray-300">
                    <span className="w-12 h-12 bg-blue-100 dark:bg-blue-900/30 rounded-full flex items-center justify-center text-blue-600">
                        📧
                    </span>
                    <div>
                      <p className="text-sm text-gray-500">Email</p>
                      <p className="font-medium">seu.email@exemplo.com</p>
                    </div>
                  </div>
                </div>
              </div>

              <ContactForm />

            </div>
          </div>
        </section>

        <footer className="py-8 text-center text-gray-500 text-sm">
          <p>© 2026 Ciro Oliveira. Todos os direitos reservados.</p>
          <p className="mt-2">
            Curioso sobre a stack? <Link href="/architecture" className="text-blue-600 hover:underline">Veja como este site foi feito.</Link>
          </p>
        </footer>
      </main>
  );
}