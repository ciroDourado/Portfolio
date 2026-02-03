<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasUuids;
use Illuminate\Database\Eloquent\Model;

class Profile extends Model
{
    use HasUuids;

    protected $fillable = [
        'name',
        'headline',
        'bio',
        'avatar_url',
        'github_url',
        'linkedin_url',
        'cv_url',
    ];

    public function uniqueIds(): array
    {
        return ['uuid'];
    }
}
