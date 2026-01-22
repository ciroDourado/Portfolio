<?php

namespace Database\Factories;

use App\Models\Category;
use Illuminate\Database\Eloquent\Factories\Factory;
use Illuminate\Support\Str;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\Skill>
 */
class SkillFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        $name = fake()->unique()->word();

        return [
            'uuid' => fake()->uuid(),
            'skill_category_id' => Category::factory(),
            'name' => ucfirst($name),
            'slug' => Str::slug($name),
            'icon_url' => null,
            'proficiency_level' => fake()->numberBetween(1, 5),
            'is_highlight' => fake()->boolean(20),
        ];
    }
}
