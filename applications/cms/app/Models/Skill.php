<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasUuids;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\BelongsToMany;

class Skill extends Model
{
    /** @use HasFactory<\Database\Factories\SkillFactory> */
    use HasFactory, HasUuids;

    protected $fillable = [
        'skill_category_id',
        'name',
        'slug',
        'icon_url',
        'proficiency_level',
        'is_highlight'
    ];

    protected $casts = [
        'is_highlight' => 'boolean',
        'proficiency_level' => 'integer',
    ];

    public function category(): BelongsTo
    {
        return $this->belongsTo(Category::class, 'skill_category_id');
    }

    public function projects(): BelongsToMany
    {
        return $this->belongsToMany(Project::class, 'project_skill');
    }

    public function uniqueIds(): array
    {
        return ['uuid'];
    }
}
