export interface Skill {
    id: string;
    name: string;
}

export interface Project {
    id: string;
    title: string;
    slug: string;
    short_description: string;
    content: string;
    repository_url: string;
    live_url: string;
    thumbnail_url: string;
    skills: Skill[];
    started_at?: string;
    finished_at?: string;
}

export interface Experience {
    id: string;
    company_name: string;
    role: string;
    description: string;
    start_date: string;
    end_date: string | null;
}