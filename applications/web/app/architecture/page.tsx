import Link from 'next/link';
import {
    FaReact, FaDocker, FaAws, FaLinux, FaDiscord
} from 'react-icons/fa';
import {
    SiGo, SiLaravel, SiPostgresql, SiRedis, SiTailwindcss, SiTypescript, SiNginx
} from 'react-icons/si';

export const metadata = {
    title: 'Arquitetura do Sistema | Ciro Oliveira',
    description: 'Deep dive na engenharia por trás deste portfólio: Go, Microservices e Cloud.',
};

export default function ArchitecturePage() {
    return (
        <main className="min-h-screen bg-gray-50 dark:bg-gray-900 pb-20">

            <section className="bg-white dark:bg-gray-800 py-20 border-b border-gray-200 dark:border-gray-700">
                <div className="max-w-4xl mx-auto px-4 text-center">
                    <h1 className="text-4xl md:text-5xl font-extrabold text-gray-900 dark:text-white mb-6">
                        Como este portfólio foi construído?
                    </h1>
                    <p className="text-xl text-gray-600 dark:text-gray-300 leading-relaxed">
                        Isto não é apenas um site estático. É um sistema distribuído
                        focado em performance, desacoplamento e automação.
                    </p>
                </div>
            </section>

            <div className="max-w-5xl mx-auto px-4 mt-16 space-y-24">

                <section>
                    <div className="flex items-center gap-4 mb-8">
                        <div className="h-10 w-1 bg-blue-600 rounded-full"></div>
                        <h2 className="text-3xl font-bold text-gray-900 dark:text-white">Visão Geral da Arquitetura</h2>
                    </div>

                    <div className="bg-white dark:bg-gray-800 p-4 rounded-2xl shadow-sm border border-gray-200 dark:border-gray-700 overflow-hidden">
                        <img
                            src="/architecture-diagram.svg"
                            alt="Diagrama de Arquitetura do Sistema"
                            className="w-full h-auto dark:opacity-90"
                        />
                    </div>
                    <p className="mt-4 text-gray-500 text-sm text-center">
                        Diagrama gerado via Mermaid.js representando o fluxo de dados em produção.
                    </p>
                </section>

                <section className="grid md:grid-cols-2 gap-12">
                    <div className="space-y-6">
                        <h3 className="text-2xl font-bold text-gray-900 dark:text-white">Backend Híbrido</h3>
                        <p className="text-gray-600 dark:text-gray-400 leading-relaxed">
                            Utilizei uma abordagem <strong>Headless CMS</strong>. O <span className="text-red-500 font-bold">Laravel Filament</span> gerencia o conteúdo (BD) pela sua velocidade de desenvolvimento, enquanto uma <span className="text-blue-500 font-bold">API em Go</span> serve os dados publicamente.
                        </p>
                        <p className="text-gray-600 dark:text-gray-400 leading-relaxed">
                            Isso garante que o painel administrativo tenha a robustez do PHP, mas o site final tenha a performance bruta e o baixo consumo de memória do Go.
                        </p>
                    </div>
                    <div className="space-y-6">
                        <h3 className="text-2xl font-bold text-gray-900 dark:text-white">Mensageria & Resiliência</h3>
                        <p className="text-gray-600 dark:text-gray-400 leading-relaxed">
                            Para o formulário de contato, implementei uma arquitetura assíncrona usando <strong>Redis</strong> e <strong>Go Workers</strong>.
                        </p>
                        <p className="text-gray-600 dark:text-gray-400 leading-relaxed">
                            A API apenas enfileira a solicitação. Um worker separado processa o envio para o Discord e Telegram. Isso garante que, mesmo sob carga alta ou falha de rede externa, nenhuma mensagem de contato seja perdida.
                        </p>
                    </div>
                </section>

                <section>
                    <h2 className="text-3xl font-bold text-gray-900 dark:text-white mb-10 text-center">Tech Stack</h2>

                    <div className="grid grid-cols-1 md:grid-cols-4 gap-6">

                        <StackCard title="Frontend" color="border-blue-500">
                            <TechItem name="Next.js 14" desc="App Router & SSR" />
                            <TechItem name="Tailwind CSS" desc="Styling & Typography" />
                            <TechItem name="TypeScript" desc="Type Safety" />
                        </StackCard>

                        <StackCard title="Core Backend" color="border-cyan-500">
                            <TechItem name="Go (Golang)" desc="High Perf. API" />
                            <TechItem name="Chi Router" desc="Lightweight HTTP" />
                            <TechItem name="Asynq" desc="Redis Job Queue" />
                        </StackCard>

                        <StackCard title="Data & CMS" color="border-purple-500">
                            <TechItem name="Laravel" desc="Headless CMS" />
                            <TechItem name="PostgreSQL" desc="Relational DB" />
                            <TechItem name="Redis" desc="Cache & Queue Broker" />
                        </StackCard>

                        <StackCard title="Infraestrutura" color="border-orange-500">
                            <TechItem name="Docker" desc="Containerization" />
                            <TechItem name="Oracle OCI" desc="Object Storage (S3)" />
                            <TechItem name="Caddy" desc="Reverse Proxy & SSL" />
                            <TechItem name="GitHub Actions" desc="CI/CD Pipeline" />
                        </StackCard>

                    </div>
                </section>

                <section className="text-center py-10">
                    <Link
                        href="https://github.com/cirodourado/portfolio"
                        target="_blank"
                        className="inline-flex items-center px-8 py-4 bg-gray-900 text-white rounded-full font-bold hover:bg-gray-800 transition shadow-lg"
                    >
                        Ver Código Fonte no GitHub
                    </Link>
                </section>

            </div>
        </main>
    );
}

function StackCard({ title, children, color }: { title: string, children: React.ReactNode, color: string }) {
    return (
        <div className={`bg-white dark:bg-gray-800 p-6 rounded-xl shadow-sm border-t-4 ${color}`}>
            <h4 className="font-bold text-lg mb-4 text-gray-900 dark:text-white">{title}</h4>
            <ul className="space-y-3">
                {children}
            </ul>
        </div>
    )
}

function TechItem({ name, desc }: { name: string, desc: string }) {
    return (
        <li className="flex justify-between items-center text-sm">
            <span className="font-medium text-gray-700 dark:text-gray-200">{name}</span>
            <span className="text-gray-400 text-xs">{desc}</span>
        </li>
    )
}