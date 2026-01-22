<?php

namespace Database\Factories;

use Illuminate\Database\Eloquent\Factories\Factory;
use Illuminate\Support\Str;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\Project>
 */
class ProjectFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        $title = fake()->sentence(3);
        $startedAt = fake()->dateTimeBetween('-2 years', '-1 year');

        return [
            'uuid' => fake()->uuid(),
            'title' => trim($title, '.'),
            'slug' => Str::slug($title),
            'short_description' => fake()->text(100),
            'content' => "# Sobre o Projeto\n\n" . fake()->paragraphs(3, true),
            'repository_url' => null,
            'live_url' => fake()->url(),
            'thumbnail_url' => null,
            'started_at' => $startedAt,
            'finished_at' => fake()->boolean(80) ? fake()->dateTimeBetween($startedAt, 'now') : null,
            'is_featured' => fake()->boolean(10),
        ];
    }
}
