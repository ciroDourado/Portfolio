export interface Skill {
    id: string;
    name: string;
}

export interface Project {
    id: string;
    title: string;
    slug: string;
    short_description: string;
    thumbnail_url: string;
    repository_url: string;
    live_url: string;
    skills: Skill[];
}