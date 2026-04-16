import * as Icons from 'react-icons/fa';
import { SoftSkill } from '@/types';

interface BeyondCodeProps {
    skills: SoftSkill[];
}

export default function BeyondCodeSection({ skills }: BeyondCodeProps) {
    if (skills.length === 0) return null;

    return (
        <section className="py-20 bg-white dark:bg-gray-800 border-t border-gray-100 dark:border-gray-800">
            <div className="max-w-7xl mx-auto px-4 md:px-8">
                <div className="text-center mb-16">
                    <h2 className="text-3xl md:text-4xl font-bold text-gray-900 dark:text-white mb-4">
                        Muito Além do Terminal
                    </h2>
                    <p className="text-lg text-gray-600 dark:text-gray-400 max-w-2xl mx-auto">
                        Acredito que o que fazemos longe do teclado molda a forma como resolvemos problemas na frente dele.
                    </p>
                </div>

                <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
                    {skills.map((skill: SoftSkill) => {
                        const iconName = skill.icon as keyof typeof Icons;
                        const IconComponent = Icons[iconName] || Icons.FaQuestion;

                        return (
                            <div key={skill.id} className="bg-gray-50 dark:bg-gray-900 p-8 rounded-2xl border border-gray-100 dark:border-gray-700 hover:-translate-y-1 transition-all duration-300">
                                <div className="w-14 h-14 bg-blue-100 dark:bg-blue-900/50 text-blue-600 rounded-xl flex items-center justify-center text-2xl mb-6">
                                    <IconComponent />
                                </div>
                                <h3 className="text-xl font-bold text-gray-900 dark:text-white mb-3">
                                    {skill.title}
                                </h3>
                                <p className="text-gray-600 dark:text-gray-400 leading-relaxed">
                                    {skill.description}
                                </p>
                            </div>
                        );
                    })}
                </div>
            </div>
        </section>
    );
}