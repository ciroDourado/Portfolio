<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasUuids;
use Illuminate\Database\Eloquent\Model;

class SoftSkill extends Model
{
    use HasUuids;

    protected $fillable = [
        'title',
        'icon',
        'description',
        'order',
    ];

    public function uniqueIds(): array
    {
        return ['uuid'];
    }
}
