<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasUuids;
use Illuminate\Database\Eloquent\Model;

class ContactMessage extends Model
{
    use HasUuids;

    protected $fillable = [
        'name',
        'email',
        'message',
        'status',
    ];

    public function uniqueIds(): array
    {
        return ['uuid'];
    }
}
