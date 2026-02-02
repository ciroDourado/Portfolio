<?php

namespace App\Filament\Resources\Projects\Schemas;

use Filament\Infolists\Components\IconEntry;
use Filament\Infolists\Components\TextEntry;
use Filament\Schemas\Schema;

class ProjectInfolist
{
    public static function configure(Schema $schema): Schema
    {
        return $schema
            ->components([
                TextEntry::make('uuid')
                    ->label('UUID'),
                TextEntry::make('title'),
                TextEntry::make('slug'),
                TextEntry::make('short_description')
                    ->placeholder('-'),
                TextEntry::make('content')
                    ->placeholder('-')
                    ->columnSpanFull(),
                TextEntry::make('repository_url')
                    ->placeholder('-'),
                TextEntry::make('live_url')
                    ->placeholder('-'),
                TextEntry::make('thumbnail_url')
                    ->placeholder('-'),
                TextEntry::make('started_at')
                    ->date()
                    ->placeholder('-'),
                TextEntry::make('finished_at')
                    ->date()
                    ->placeholder('-'),
                IconEntry::make('is_featured')
                    ->boolean(),
                TextEntry::make('created_at')
                    ->dateTime()
                    ->placeholder('-'),
                TextEntry::make('updated_at')
                    ->dateTime()
                    ->placeholder('-'),
            ]);
    }
}
