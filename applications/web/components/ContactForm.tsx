'use client';

import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import toast, { Toaster } from 'react-hot-toast';
import { FaPaperPlane, FaSpinner } from 'react-icons/fa';

const contactSchema = z.object({
    name: z.string().min(2, 'O nome precisa ter pelo menos 2 caracteres'),
    email: z.string().email('Digite um e-mail válido'),
    message: z.string().min(10, 'Sua mensagem deve ter pelo menos 10 caracteres'),
});

type ContactFormData = z.infer<typeof contactSchema>;

export default function ContactForm() {
    const [isLoading, setIsLoading] = useState(false);

    const {
        register,
        handleSubmit,
        reset,
        formState: { errors },
    } = useForm<ContactFormData>({
        resolver: zodResolver(contactSchema),
    });

    const onSubmit = async (data: ContactFormData) => {
        setIsLoading(true);

        try {
            const apiUrl = process.env.NEXT_PUBLIC_API_URL;

            const response = await fetch(`${apiUrl}/contact`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
            });

            if (!response.ok) throw new Error('Falha ao enviar');

            toast.success('Mensagem recebida! Logo entrarei em contato.');
            reset();
        } catch (error) {
            console.error(error);
            toast.error('Ocorreu um erro ao enviar. Tente novamente.');
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="w-full max-w-lg mx-auto bg-white dark:bg-gray-800 p-8 rounded-2xl shadow-lg border border-gray-100 dark:border-gray-700">
            <Toaster position="bottom-right" />

            <h3 className="text-2xl font-bold text-gray-900 dark:text-white mb-2">
                Vamos conversar?
            </h3>
            <p className="text-gray-500 dark:text-gray-400 mb-6">
                Tem um projeto em mente ou quer apenas trocar uma ideia? Me mande uma mensagem!
            </p>

            <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">

                <div>
                    <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                        Seu Nome
                    </label>
                    <input
                        {...register('name')}
                        type="text"
                        className={`w-full px-4 py-3 rounded-lg bg-gray-50 dark:bg-gray-900 border ${
                            errors.name ? 'border-red-500' : 'border-gray-200 dark:border-gray-700'
                        } focus:ring-2 focus:ring-blue-500 outline-none transition-all dark:text-white`}
                    />
                    {errors.name && (
                        <p className="text-red-500 text-xs mt-1">{errors.name.message}</p>
                    )}
                </div>

                <div>
                    <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                        Seu E-mail
                    </label>
                    <input
                        {...register('email')}
                        type="email"
                        className={`w-full px-4 py-3 rounded-lg bg-gray-50 dark:bg-gray-900 border ${
                            errors.email ? 'border-red-500' : 'border-gray-200 dark:border-gray-700'
                        } focus:ring-2 focus:ring-blue-500 outline-none transition-all dark:text-white`}
                    />
                    {errors.email && (
                        <p className="text-red-500 text-xs mt-1">{errors.email.message}</p>
                    )}
                </div>

                <div>
                    <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                        Mensagem
                    </label>
                    <textarea
                        {...register('message')}
                        rows={4}
                        className={`w-full px-4 py-3 rounded-lg bg-gray-50 dark:bg-gray-900 border ${
                            errors.message ? 'border-red-500' : 'border-gray-200 dark:border-gray-700'
                        } focus:ring-2 focus:ring-blue-500 outline-none transition-all resize-none dark:text-white`}
                        placeholder="Gostaria de falar sobre..."
                    />
                    {errors.message && (
                        <p className="text-red-500 text-xs mt-1">{errors.message.message}</p>
                    )}
                </div>

                <button
                    type="submit"
                    disabled={isLoading}
                    className="w-full flex items-center justify-center gap-2 bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 rounded-lg transition-all disabled:opacity-70 disabled:cursor-not-allowed mt-2"
                >
                    {isLoading ? (
                        <>
                            <FaSpinner className="animate-spin" /> Enviando...
                        </>
                    ) : (
                        <>
                            <FaPaperPlane /> Enviar Mensagem
                        </>
                    )}
                </button>
            </form>
        </div>
    );
}