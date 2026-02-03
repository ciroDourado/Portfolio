import { Experience } from '@/types';

const formatDate = (dateString: string | null) => {
    if (!dateString) return 'Atualmente';
    return new Date(dateString).toLocaleDateString('pt-BR', { month: 'short', year: 'numeric' });
};

export default function ExperienceTimeline({ experiences }: { experiences: Experience[] }) {
    return (
        <div className="relative border-l border-gray-200 dark:border-gray-700 ml-3 md:ml-6 space-y-10">
            {experiences.map((exp: Experience) => (
                <div key={exp.id} className="relative pl-8 md:pl-12 group">

                    <div className="absolute -left-1.5 top-1.5 h-3 w-3 rounded-full border border-white bg-gray-300 dark:border-gray-900 dark:bg-gray-600 group-hover:bg-blue-600 transition-colors" />

                    <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-1">
                        <h3 className="text-lg font-bold text-gray-900 dark:text-white">
                            {exp.role} <span className="text-gray-500 font-normal">at {exp.company_name}</span>
                        </h3>
                        <time className="text-sm font-mono text-gray-500 dark:text-gray-400">
                            {formatDate(exp.start_date)} — {formatDate(exp.end_date)}
                        </time>
                    </div>

                    <div
                        className="prose prose-sm dark:prose-invert text-gray-600 dark:text-gray-300 max-w-none mt-2"
                        dangerouslySetInnerHTML={{ __html: exp.description }}
                    />
                </div>
            ))}
        </div>
    );
}