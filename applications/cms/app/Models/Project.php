<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasUuids;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsToMany;

class Project extends Model
{
    /** @use HasFactory<\Database\Factories\ProjectFactory> */
    use HasFactory, HasUuids;

    protected $fillable = [
        'title',
        'slug',
        'short_description',
        'content',
        'repository_url',
        'live_url',
        'thumbnail_url',
        'started_at',
        'finished_at',
        'is_featured'
    ];

    protected $casts = [
        'started_at' => 'date',
        'finished_at' => 'date',
        'is_featured' => 'boolean',
    ];

    public function skills(): BelongsToMany
    {
        return $this->belongsToMany(Skill::class, 'project_skill');
    }

    public function uniqueIds(): array
    {
        return ['uuid'];
    }
}
