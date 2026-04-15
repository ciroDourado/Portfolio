import { fetchAPI } from '@/lib/api';
import {Experience, Skill, SkillCategory} from '@/types';

export default async function SkillsSection({ categories }: { categories: SkillCategory[] }) {
    return (
        <section className="py-20 bg-white dark:bg-gray-800">
            <div className="max-w-7xl mx-auto px-4 md:px-8">
                <h2 className="text-3xl md:text-4xl font-bold mb-12 text-center text-gray-900 dark:text-white">
                    Minha Tech Stack
                </h2>

                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
                    {categories.map((category: SkillCategory) => (
                        <div
                            key={`category_${category.id}`}
                            className="p-6 rounded-2xl bg-gray-50 dark:bg-gray-900 border border-gray-100 dark:border-gray-700 shadow-sm"
                        >
                            <h3 className="text-xl font-bold mb-6 text-gray-800 dark:text-gray-200 border-b pb-2 border-gray-200 dark:border-gray-700">
                                {category.name}
                            </h3>

                            <div className="flex flex-wrap gap-3">
                                {category.skills?.map((skill: Skill) => (
                                    <div
                                        key={`skill_${skill.id}`}
                                        className="flex items-center gap-2 px-3 py-2 bg-white dark:bg-gray-800 rounded-lg shadow-sm border border-gray-100 dark:border-gray-700 hover:scale-105 transition-transform"
                                    >
                                        {skill.icon_url && (
                                            <img
                                                src={skill.icon_url}
                                                alt={skill.name}
                                                className="w-5 h-5 object-contain"
                                            />
                                        )}
                                        <span className="text-sm font-medium text-gray-700 dark:text-gray-300">
                                            {skill.name}
                                        </span>
                                    </div>
                                ))}
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </section>
    );
}