<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasUuids;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\HasMany;

class Category extends Model
{
    /** @use HasFactory<\Database\Factories\CategoryFactory> */
    use HasFactory, HasUuids;

    protected $table = 'skill_categories';

    protected $fillable = [
        'name',
        'slug',
        'order',
    ];

    public function skills(): HasMany
    {
        return $this->hasMany(Skill::class, 'skill_category_id');
    }

    public function uniqueIds(): array
    {
        return ['uuid'];
    }
}
